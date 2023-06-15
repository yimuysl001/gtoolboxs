package automatic

//	type program struct {
//		log service.Logger
//		cfg *service.Config
//		f   func()
//	}
//
//	func (p *program) Start(s service.Service) error {
//		go p.run(p.f)
//		return nil
//	}
//
// func (p *program) run(f func()) {
//
//		f()
//		//这里写运行时的代码
//		wg.Done()
//	}
//
//	func (p *program) Stop(s service.Service) error {
//		return nil
//	}
//
// var wg sync.WaitGroup
//
//	func Auto(f func()) {
//		wg.Add(1)
//		svcConfig := &service.Config{
//			Name:        sys.GetSysName(),
//			DisplayName: sys.GetSysName(),
//			Description: sys.GetSysName(),
//		}
//		prg := &program{f: f}
//		s, err := service.New(prg, svcConfig)
//		if err != nil {
//			fmt.Println(err.Error())
//		}
//		if len(os.Args) > 1 {
//			if os.Args[1] == "install" {
//				x := s.Install()
//				if x != nil {
//					fmt.Println("error:", x.Error())
//					return
//				}
//				fmt.Println("服务安装成功")
//				return
//			} else if os.Args[1] == "uninstall" {
//				x := s.Uninstall()
//				if x != nil {
//					fmt.Println("error:", x.Error())
//					return
//				}
//				fmt.Println("服务卸载成功")
//				return
//			}
//		}
//
//		err = s.Run()
//		if err != nil {
//			fmt.Println(err.Error())
//		}
//		wg.Wait()
//	}
//
//	桌面快捷方式
const startup = "%s\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\%s.lnk"

// 开机自启动快捷方式
const desktopPth = "%s\\Desktop\\%s.lnk"
