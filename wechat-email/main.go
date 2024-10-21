//package main
//
//import (
//	"gopkg.in/gomail.v2"
//	"log"
//	"os"
//)
//
//func main() {
//	args := os.Args //获取用户输入的所有参数
//	if args == nil {
//		return
//	}
//	message := "hello i am everai"
//	m := gomail.NewMessage()
//	m.SetHeader("From", "expvent@expvent.com")
//	m.SetHeader("To", "1835783944@qq.com")
//	m.SetHeader("Subject", "邮件主题")
//	m.SetBody("text/html", message)
//
//	d := gomail.NewPlainDialer("smtp.exmail.qq.com", 465, "expvent@expvent.com", "jxqFKUG8q8PYEify") // 发送邮件服务器、端口、发件人账号、发件人密码
//	if err := d.DialAndSend(m); err != nil {
//		log.Println("发送失败", err)
//		return
//	}
//
//	log.Println("done.发送成功")
//}

package main

import (
	"gopkg.in/gomail.v2"
	"log"
)

func main() {
	//// 获取用户输入的所有参数
	//args := os.Args
	//if len(args) < 2 {
	//	log.Println("请提供目标邮箱")
	//	return
	//}

	recipient := "1835783944@qq.com" // 目标邮箱
	message := `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>割接通知</title>
    <style>
        body {
            display: flex;
            justify-content: center;
            align-items: center;
            min-height: 100vh; /* 确保内容在页面中垂直居中 */
            text-align: center; /* 默认文本居中 */
            margin: 0; /* 去掉默认的边距 */
            background-color: #f9f9f9; /* 可选：设置背景色 */
        }
        .content {
            max-width: 600px; /* 设置最大宽度 */
            padding: 20px; /* 内边距 */
            background-color: #fff; /* 可选：设置内容背景 */
            border-radius: 8px; /* 可选：圆角效果 */
            box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1); /* 可选：阴影效果 */
        }
        .right-align {
            text-align: right; /* 右对齐 */
        }
        .left-align {
            text-align: left; /* 右对齐 */
        }
        .indent {
    		padding-left: 2em; /* 根据字体大小设置合适的缩进 */
		}
        .header-cs {
            height: 100px;
      }
    </style>
</head>
<body>
  <div>
    <header class="header-cs"> </header>
    <div class="content">
      <div >
        <h2>2024年08月28日01:00~05:00 上海二可用区B传输线路优化割接</h2>
        <hr style="border: 1px solid #ccc; margin: 10px 0;"> <!-- 添加分割横线 -->
      </div>
        <div class="left-align">
          <p>尊敬的EverAI用户，您好！</p>
          <p class="indent">接供应商割接通知，计划对上海二可用区B传输线路进行割接</p>
          <p><strong>割接时间：</strong>北京时间2024年08月28日01:00~05:00</p>
          <p><strong>割接影响：</strong>割接期间，上海二可用区B的内网互访会出现1~2次秒级网络抖动，延迟变化1ms以内。</p>
          <p>如有任何疑问，请联系技术支持 <strong>4000188113</strong>。</p>        
        </div>
        <div class= "right-align">
          <p>EverAI云计算团队</p>
          <p>2024-08-21</p>
    	</div>
    </div>
    <footer>
      <p>访问<a href="https://everai.expvent.com/">EverAI官网</a></p>
      <p>Copyright © 2022-2024 Expvent 上海幂进信息科技有限公司</p>
    </footer>
  </div>
    
</body>
  
</html>`

	m := gomail.NewMessage()
	m.SetHeader("From", "expvent@expvent.com") // 设置发件人名称
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", "上海二可用区B传输线路优化割接")
	m.SetBody("text/html", message)

	d := gomail.NewPlainDialer("smtp.exmail.qq.com", 465, "expvent@expvent.com", "jxqFKUG8q8PYEify") // 发送邮件服务器、端口、发件人账号、发件人密码
	if err := d.DialAndSend(m); err != nil {
		log.Println("发送失败:", err)
		return
	}

	log.Println("邮件发送成功")
}
