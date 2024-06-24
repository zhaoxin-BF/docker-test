#!/bin/bash

# 获取命令行参数
source_dir=$1
target_dir=$2

# 创建目标目录
mkdir -p "$target_dir"

# 遍历源文件目录下的所有 .service 文件
for file in "$source_dir"/*.service; do
    # 检查文件是否存在
    if [ -f "$file" ]; then
        # 获取文件名
        filename=$(basename "$file")
        # 构建目标文件路径
        target_file="$target_dir/$filename"

        # 读取源文件内容
        content=$(cat "$file")

        # 替换 workerSpace 字段
        new_content="${content/workerSpace=/workerSpace=\/my\/new\/path}"

        # 将替换后的内容写入目标文件
        echo "$new_content" > "$target_file"

        echo "File $filename processed successfully."
    fi
done

echo "All files processed."