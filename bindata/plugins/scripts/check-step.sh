#!/usr/bin/env bash

set -e
set -o pipefail

installGuide() {
  model="$1"
  case "$model" in
      "Docker")
          echo "Docker Installation Guide:"
          echo "1、Use \"apt install docker.io\" for installation, or refer to the official Docker documentation for installation. You need to use Docker version > 24.0.0."
          echo "  - docker documentation：https://docs.docker.com/engine/install/ubuntu/"
          echo "2、If this is a GPU Host, Please refer to the following documents to go to Docker GPURuntime"
          echo "  - documentation：https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/latest/install-guide.html"
          ;;
      "GPUDriver")
          echo "GPUDriver Installation Guide:"
          echo "Option 1: Automatic installation using the standard Ubuntu repository"
          echo "ubuntu-drivers devices"
          echo "sudo ubuntu-drivers autoinstall"
          echo "Option 2: Manual installation using the official NVIDIA driver"
          echo "Download the GPU driver from the NVIDIA website: https://www.nvidia.com/Download/index.aspx"
          echo "wget -c http://us.download.nvidia.com/XFree86/Linux-x86_64/535.**/NVIDIA-Linux-x86_64-535.**.run"
          echo "sudo chmod a+x NVIDIA-Linux-x86_64-440.82.run"
          echo "sudo ./NVIDIA-Linux-x86_64-440.82.run -no-opengl-files -no-nouveau-check"
          ;;
      "DCGM")
          # commands
          echo "DCGM Installation Guide:"
          echo "Please refer to the official installation documentation for the required version 3.3.5 or above."
          echo "https://developer.nvidia.com/dcgm#Downloads"
          ;;
      "GLIBC")
          # commands
          echo "GLIBC Installation Guide:"
          echo "1. Open the /etc/apt/sources.list file:"
          echo "   vi /etc/apt/sources.list"
          echo "2. Add one of the following repository entries:"
          echo "   - deb http://security.ubuntu.com/ubuntu jammy-security main"
          echo "   - deb http://th.archive.ubuntu.com/ubuntu jammy main"
          echo "3. Update the package lists:"
          echo "   apt update"
          echo "4. Upgrade the installed packages:"
          echo "   apt upgrade"
          echo "5. Remove any unused packages:"
          echo "   apt autoremove"
          echo "6. Install the GLIBC package:"
          echo "   apt install libc6"
          echo "7. Check the GLIBC version by running:"
          echo "   strings /lib/x86_64-linux-gnu/libc.so.6 | grep GLIBC"
          echo "   The version should be greater than 2.35."
          ;;
      "DockerGPURuntime")
          # commands
          echo "DockerGPURuntime Installation Guide:"
          echo "1. Before installing the docker GPU runtime, please check if docker.io is installed. If not, please refer to: https://docs.docker.com/engine/install/ubuntu/"
          echo "2. For the DockerGPURuntime installation method, please refer to the NVIDIA official documentation, https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/latest/install-guide.html"
          echo "3. After the installation is complete and docker is restarted, please run the following commands to check if the GPU runtime is successfully installed"
          echo "   - Start a container that uses GPU: docker run -d -name test-gpu-nginx --rm --gpus all nginx"
          echo "   - Execute the command to check: docker exec -it test-gpu-nginx nvidia-smi"
          echo "correct step 3 printing of the GPU info means the installation is complete"
         ;;
      *)
         # default commands
         echo "EverAI-Resource-Node Installation Guide:"
         ;;
  esac
}

checkDocker() {
  # Check if Docker is installed
  if command -v docker &> /dev/null; then
      echo "Docker is installed."
  else
      echo "[ERROR]: Docker is not installed."
      installGuide Docker
      return 1
  fi

  # Check the Docker version
  local docker_version=$(docker --version | awk '{print $3}')
  echo "Docker version: $docker_version"

  # Check if the version is compatible with the required version
  local required_version="26.0.0"
  if [[ "$docker_version" < "$required_version" ]]; then
      echo "[ERROR]: Docker version $docker_version is not compatible. Required version is at least $required_version."
      installGuide Docker
      return 1
  fi
  return 0
}

checkGPUDriver() {
  # Check if NVIDIA driver is installed
  if command -v nvidia-smi &> /dev/null; then
    echo "NVIDIA driver is installed."
  else
    echo "NVIDIA driver is not installed."
    installGuide GPUDriver
    return 1
  fi

  # Check the NVIDIA driver version
  local driver_version=$(nvidia-smi -h | head -n 1 | awk '{print$6}' | awk -F 'v' '{print$2}')
  echo "NVIDIA driver version: $driver_version"

  # Check if the driver version is compatible with the required version
  local required_version="535.00"
  if [[ "$driver_version" < "$required_version" ]]; then
    echo "NVIDIA driver version $driver_version is not compatible. Required version is at least $required_version."
    installGuide GPUDriver
    return 1
  fi

  return 0
}

checkDCGM() {
  # Check if DCGM is installed
  if command -v dcgmi &> /dev/null; then
    echo "DCGM is installed."
  else
    echo "DCGM is not installed."
    installGuide DCGM
    return 1
  fi

  # Check the DCGM version
  local dcgm_version=$(dcgmi -v | grep "Version"| awk '{print $3}')
  echo "DCGM version: $dcgm_version"

  # Check if the DCGM version is compatible with the required version
  local required_version="3.3.5"
  if [[ "$dcgm_version" < "$required_version" ]]; then
    echo "DCGM version $dcgm_version is not compatible. Required version is at least $required_version."
    installGuide DCGM
    return 1
  fi

  return 0
}

checkGLIBC() {
  # Check the GLIBC version
  local glibc_version=$(strings /lib/x86_64-linux-gnu/libc.so.6 |grep GLIBC_|tail -n 2|head -n 1|awk -F '_' '{print $2}')
  echo "GLIBC version: $glibc_version"

  # Check if the GLIBC  version is compatible with the required version
  local required_version="2.35"
  if [[ "$glibc_version" < "$required_version" ]]; then
    echo "GLIBC  version $glibc_version is not compatible. Required version is at least $required_version"
    installGuide GLIBC
    return 1
  fi
}

checkDockerGPURuntime() {
#  if dpkg-query -W -f='${Status}' nvidia-container-toolkit 2>/dev/null | grep -q "install ok installed"; then
#      echo "nvidia-container-toolkit is already installed."
#  else
#      echo "nvidia-container-toolkit is not installed. Installing now..."
#      installGuide DockerGPURuntime
## sudo apt-get install -y nvidia-container-toolkit
#  fi

  TEST_NAME="this-is-container-gpus-test-nginx"
# check docker whether set nvidia-ctk runtime
  if docker run -d --rm  --name $TEST_NAME --gpus all nginx 2>/dev/null | docker exec -it $TEST_NAME nvidia-smi 2>/dev/null; then
    echo "docker nvidia-ctk runtime is already set."
    docker kill $TEST_NAME
  else
    echo "docker nvidia-ctk runtime is not set, please check"
  fi
}

function check() {
  isGPUHost="${1:-true}"

  checkDocker

  if [ "$isGPUHost" = true ]; then
    checkGPUDriver
    checkDCGM
    checkDockerGPURuntime
  elif [ "$isGPUHost" = false ]; then
    # Handle non-GPU host case
    echo "Running on a non-GPU host"
  else
    echo "Invalid value for IS_GPU_HOST. Please use 'true' or 'false'."
    exit 1
  fi

  checkGLIBC
}

# Call the check function with the resource path
if [ "$#" -ne 1 ]; then
  echo "Usage: $0 <IS_GPU_HOST>" # IS_GPU_HOST is a boolean value, default is true
  exit 1
fi

check "$1"