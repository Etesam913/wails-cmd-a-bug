## cmd+A bug

* When you open a file dialog and do cmd+a the app quits
* This happens when a go function is called from the frontend

https://github.com/Etesam913/wails-cmd-a-bug/assets/55665282/7fcf500e-a58a-4345-aac8-4bc8e0bdff1f


### The output when cmd+a is ran

```
*** Terminating app due to uncaught exception 'NSInternalInconsistencyException', reason: 'assertion failed: 'withHintInProgress has been invoked on a thread which is incompatible with AppKit; the problem is likely in a much shallower frame in the backtrace' in withHintInProgress on line 1594 of file /AppleInternal/Library/BuildRoots/4e1473ee-9f66-11ee-8daf-cedaeb4cabe2/Library/Caches/com.apple.xbs/Sources/ViewBridge/ViewBridgeUtilities.m'
*** First throw call stack:
(
        0   CoreFoundation                      0x0000000180dbc540 __exceptionPreprocess + 176
        1   libobjc.A.dylib                     0x00000001808adeb4 objc_exception_throw + 60
        2   CoreFoundation                      0x0000000180de188c _CFBundleGetValueForInfoKey + 0
        3   ViewBridge                          0x000000018942885c withHintInProgress + 260
        4   ViewBridge                          0x00000001893ecb90 -[NSRemoteView _invalidateWithException:withCleanup:] + 272
        5   ViewBridge                          0x00000001893f912c -[NSRemoteView supplementalTargetForAction:sender:] + 308
        6   AppKit                              0x00000001845f8a90 _objectFromResponderChainWhichRespondsToAction + 256
        7   AppKit                              0x00000001845f4d14 _NSTargetForSendAction + 1616
        8   AppKit                              0x00000001846c9dd4 -[NSApplication(NSResponder) sendAction:to:from:] + 92
        9   cmd-a-bug                           0x0000000104a98ad0 selectAll + 40
        10  cmd-a-bug                           0x0000000104a98a9c _cgo_97cb83dfff12_Cfunc_selectAll + 28
        11  cmd-a-bug                           0x00000001046bcf0c runtime.asmcgocall.abi0 + 124
)
libc++abi: terminating due to uncaught exception of type NSException
SIGABRT: abort
PC=0x180c320dc m=8 sigcode=0
signal arrived during cgo execution

goroutine 1171 gp=0x140001c6fc0 m=8 mp=0x14000099b08 [syscall]:
runtime.cgocall(0x104a98a80, 0x140002a57a8)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/cgocall.go:157 +0x44 fp=0x140002a5770 sp=0x140002a5730 pc=0x104652604
github.com/wailsapp/wails/v3/pkg/application._Cfunc_selectAll()
        _cgo_gotypes.go:782 +0x30 fp=0x140002a57a0 sp=0x140002a5770 pc=0x104a80e90
github.com/wailsapp/wails/v3/pkg/application.newSelectAllMenuItem.func1(0x0?)
        /Users/etesam/Coding/wails/v3/pkg/application/menuitem_darwin.go:519 +0x1c fp=0x140002a57b0 sp=0x140002a57a0 pc=0x104a8f4ac
github.com/wailsapp/wails/v3/pkg/application.(*MenuItem).handleClick.gowrap1()
        /Users/etesam/Coding/wails/v3/pkg/application/menuitem.go:223 +0x30 fp=0x140002a57d0 sp=0x140002a57b0 pc=0x104a71be0
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x140002a57d0 sp=0x140002a57d0 pc=0x1046bd114
created by github.com/wailsapp/wails/v3/pkg/application.(*MenuItem).handleClick in goroutine 1170
        /Users/etesam/Coding/wails/v3/pkg/application/menuitem.go:223 +0x108

goroutine 1 gp=0x140000021c0 m=0 mp=0x1050e8f40 [syscall, locked to thread]:
runtime.cgocall(0x104a95140, 0x1400013b8f8)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/cgocall.go:157 +0x44 fp=0x1400013b8c0 sp=0x1400013b880 pc=0x104652604
github.com/wailsapp/wails/v3/pkg/application._Cfunc_run()
        _cgo_gotypes.go:770 +0x30 fp=0x1400013b8f0 sp=0x1400013b8c0 pc=0x104a80e40
github.com/wailsapp/wails/v3/pkg/application.(*macosApp).run(0x140000ac250)
        /Users/etesam/Coding/wails/v3/pkg/application/application_darwin.go:234 +0xd4 fp=0x1400013b980 sp=0x1400013b8f0 pc=0x104a84594
github.com/wailsapp/wails/v3/pkg/application.(*App).Run(0x140001b2308)
        /Users/etesam/Coding/wails/v3/pkg/application/application.go:554 +0x398 fp=0x1400013ba10 sp=0x1400013b980 pc=0x104a66718
main.main()
        /Users/etesam/Coding/cmd-a-bug/main.go:74 +0x1e8 fp=0x1400013bf60 sp=0x1400013ba10 pc=0x104a944f8
runtime.main()
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/proc.go:271 +0x228 fp=0x1400013bfd0 sp=0x1400013bf60 pc=0x10468a958
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x1400013bfd0 sp=0x1400013bfd0 pc=0x1046bd114

goroutine 18 gp=0x1400009c380 m=nil [force gc (idle)]:
runtime.gopark(0x104ceb870, 0x1050e6770, 0x11, 0xa, 0x1)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/proc.go:402 +0xe0 fp=0x1400005e770 sp=0x1400005e740 pc=0x10468adb0
runtime.goparkunlock(0x1050e6770?, 0x0?, 0x0?, 0x0?)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/proc.go:408 +0x34 fp=0x1400005e7a0 sp=0x1400005e770 pc=0x10468ae44
runtime.forcegchelper()
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/proc.go:326 +0xb4 fp=0x1400005e7d0 sp=0x1400005e7a0 pc=0x10468abd4
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x1400005e7d0 sp=0x1400005e7d0 pc=0x1046bd114
created by runtime.init.6 in goroutine 1
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/proc.go:314 +0x24

goroutine 19 gp=0x1400009c540 m=nil [GC sweep wait]:
runtime.gopark(0x104ceb870, 0x1050e6d00, 0xc, 0x9, 0x1)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/proc.go:402 +0xe0 fp=0x1400005ef40 sp=0x1400005ef10 pc=0x10468adb0
runtime.goparkunlock(0x1050e6d00?, 0x0?, 0x0?, 0x0?)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/proc.go:408 +0x34 fp=0x1400005ef70 sp=0x1400005ef40 pc=0x10468ae44
runtime.bgsweep(0x140000aa000)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/mgcsweep.go:278 +0x9c fp=0x1400005efb0 sp=0x1400005ef70 pc=0x10467583c
runtime.gcenable.gowrap1()
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/mgc.go:203 +0x28 fp=0x1400005efd0 sp=0x1400005efb0 pc=0x10466a1e8
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x1400005efd0 sp=0x1400005efd0 pc=0x1046bd114
created by runtime.gcenable in goroutine 1
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/mgc.go:203 +0x6c

goroutine 20 gp=0x1400009c700 m=nil [GC scavenge wait]:
runtime.gopark(0x104ceb870, 0x1050e7640, 0xd, 0xa, 0x2)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/proc.go:402 +0xe0 fp=0x1400005f720 sp=0x1400005f6f0 pc=0x10468adb0
runtime.goparkunlock(0x1050e7640?, 0x8?, 0xb0?, 0x1?)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/proc.go:408 +0x34 fp=0x1400005f750 sp=0x1400005f720 pc=0x10468ae44
runtime.(*scavengerState).park(0x1050e7640)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/mgcscavenge.go:425 +0x4c fp=0x1400005f780 sp=0x1400005f750 pc=0x104672e8c
runtime.bgscavenge(0x140000aa000)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/mgcscavenge.go:653 +0x44 fp=0x1400005f7b0 sp=0x1400005f780 pc=0x1046733c4
runtime.gcenable.gowrap2()
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/mgc.go:204 +0x28 fp=0x1400005f7d0 sp=0x1400005f7b0 pc=0x10466a188
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x1400005f7d0 sp=0x1400005f7d0 pc=0x1046bd114
created by runtime.gcenable in goroutine 1
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/mgc.go:204 +0xac

goroutine 2 gp=0x140000036c0 m=nil [finalizer wait]:
runtime.gopark(0x104ceb5c0, 0x10514ebc0, 0x10, 0xa, 0x1)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/proc.go:402 +0xe0 fp=0x14000062590 sp=0x14000062560 pc=0x10468adb0
runtime.runfinq()
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/mfinal.go:194 +0xf0 fp=0x140000627d0 sp=0x14000062590 pc=0x104669380
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x140000627d0 sp=0x140000627d0 pc=0x1046bd114
created by runtime.createfing in goroutine 1
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/mfinal.go:164 +0x4c

goroutine 37 gp=0x140001c61c0 m=nil [select]:
runtime.gopark(0x104ceb8d8, 0x0, 0x9, 0x3, 0x1)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/proc.go:402 +0xe0 fp=0x14000216d60 sp=0x14000216d30 pc=0x10468adb0
runtime.selectgo(0x14000216f38, 0x14000216ef8, 0x140002020c0?, 0x0, 0x14000220270?, 0x1)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/select.go:327 +0x788 fp=0x14000216eb0 sp=0x14000216d60 pc=0x10469ba58
net/http.(*persistConn).writeLoop(0x140000e4000)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/net/http/transport.go:2444 +0xa0 fp=0x14000216fb0 sp=0x14000216eb0 pc=0x1048c3ab0
net/http.(*Transport).dialConn.gowrap3()
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/net/http/transport.go:1800 +0x28 fp=0x14000216fd0 sp=0x14000216fb0 pc=0x1048c0d38
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x14000216fd0 sp=0x14000216fd0 pc=0x1046bd114
created by net/http.(*Transport).dialConn in goroutine 9
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/net/http/transport.go:1800 +0x10e8

goroutine 36 gp=0x140001c6380 m=nil [select]:
runtime.gopark(0x104ceb8d8, 0x0, 0x9, 0x3, 0x1)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/proc.go:402 +0xe0 fp=0x14000075bf0 sp=0x14000075bc0 pc=0x10468adb0
runtime.selectgo(0x14000075e40, 0x14000075d98, 0x104aa7a67?, 0x0, 0x1055c4a01?, 0x1)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/select.go:327 +0x788 fp=0x14000075d40 sp=0x14000075bf0 pc=0x10469ba58
net/http.(*persistConn).readLoop(0x140000e4000)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/net/http/transport.go:2261 +0x618 fp=0x14000075fb0 sp=0x14000075d40 pc=0x1048c27e8
net/http.(*Transport).dialConn.gowrap2()
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/net/http/transport.go:1799 +0x28 fp=0x14000075fd0 sp=0x14000075fb0 pc=0x1048c0d98
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x14000075fd0 sp=0x14000075fd0 pc=0x1046bd114
created by net/http.(*Transport).dialConn in goroutine 9
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/net/http/transport.go:1799 +0x10a0

goroutine 8 gp=0x140001c6540 m=nil [sleep]:
runtime.gopark(0x104ceb8b0, 0x14000212000, 0x13, 0xe, 0x1)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/proc.go:402 +0xe0 fp=0x14000065730 sp=0x14000065700 pc=0x10468adb0
time.Sleep(0x3b9aca00)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/time.go:195 +0xf8 fp=0x14000065770 sp=0x14000065730 pc=0x1046b9d38
main.main.func1()
        /Users/etesam/Coding/cmd-a-bug/main.go:69 +0x44 fp=0x140000657d0 sp=0x14000065770 pc=0x104a94584
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x140000657d0 sp=0x140000657d0 pc=0x1046bd114
created by main.main in goroutine 1
        /Users/etesam/Coding/cmd-a-bug/main.go:62 +0x1e0

goroutine 50 gp=0x14000284000 m=nil [chan receive]:
runtime.gopark(0x104ceb560, 0x1400017a7d8, 0xe, 0x7, 0x2)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/proc.go:402 +0xe0 fp=0x14000064ee0 sp=0x14000064eb0 pc=0x10468adb0
runtime.chanrecv(0x1400017a780, 0x14000064fc0, 0x1)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/chan.go:583 +0x26c fp=0x14000064f60 sp=0x14000064ee0 pc=0x10465459c
runtime.chanrecv1(0x0?, 0x14000064f58?)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/chan.go:442 +0x14 fp=0x14000064f90 sp=0x14000064f60 pc=0x1046542f4
github.com/wailsapp/wails/v3/pkg/application.(*App).Run.func1()
        /Users/etesam/Coding/wails/v3/pkg/application/application.go:496 +0x44 fp=0x14000064fd0 sp=0x14000064f90 pc=0x104a66ec4
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x14000064fd0 sp=0x14000064fd0 pc=0x1046bd114
created by github.com/wailsapp/wails/v3/pkg/application.(*App).Run in goroutine 1
        /Users/etesam/Coding/wails/v3/pkg/application/application.go:494 +0xc0

goroutine 51 gp=0x140002841c0 m=nil [chan receive]:
runtime.gopark(0x104ceb560, 0x1400017a838, 0xe, 0x7, 0x2)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/proc.go:402 +0xe0 fp=0x1400020a6e0 sp=0x1400020a6b0 pc=0x10468adb0
runtime.chanrecv(0x1400017a7e0, 0x1400020a7c0, 0x1)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/chan.go:583 +0x26c fp=0x1400020a760 sp=0x1400020a6e0 pc=0x10465459c
runtime.chanrecv1(0x0?, 0x0?)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/chan.go:442 +0x14 fp=0x1400020a790 sp=0x1400020a760 pc=0x1046542f4
github.com/wailsapp/wails/v3/pkg/application.(*App).Run.func2()
        /Users/etesam/Coding/wails/v3/pkg/application/application.go:502 +0x44 fp=0x1400020a7d0 sp=0x1400020a790 pc=0x104a66dc4
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x1400020a7d0 sp=0x1400020a7d0 pc=0x1046bd114
created by github.com/wailsapp/wails/v3/pkg/application.(*App).Run in goroutine 1
        /Users/etesam/Coding/wails/v3/pkg/application/application.go:500 +0x100

goroutine 52 gp=0x14000284380 m=nil [chan receive]:
runtime.gopark(0x104ceb560, 0x1400017a778, 0xe, 0x7, 0x2)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/proc.go:402 +0xe0 fp=0x1400020aee0 sp=0x1400020aeb0 pc=0x10468adb0
runtime.chanrecv(0x1400017a720, 0x1400020afc0, 0x1)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/chan.go:583 +0x26c fp=0x1400020af60 sp=0x1400020aee0 pc=0x10465459c
runtime.chanrecv1(0x0?, 0x0?)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/chan.go:442 +0x14 fp=0x1400020af90 sp=0x1400020af60 pc=0x1046542f4
github.com/wailsapp/wails/v3/pkg/application.(*App).Run.func3()
        /Users/etesam/Coding/wails/v3/pkg/application/application.go:508 +0x44 fp=0x1400020afd0 sp=0x1400020af90 pc=0x104a66cc4
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x1400020afd0 sp=0x1400020afd0 pc=0x1046bd114
created by github.com/wailsapp/wails/v3/pkg/application.(*App).Run in goroutine 1
        /Users/etesam/Coding/wails/v3/pkg/application/application.go:506 +0x140

goroutine 53 gp=0x14000284540 m=nil [chan receive]:
runtime.gopark(0x104ceb560, 0x1400017a658, 0xe, 0x7, 0x2)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/proc.go:402 +0xe0 fp=0x1400020b6e0 sp=0x1400020b6b0 pc=0x10468adb0
runtime.chanrecv(0x1400017a600, 0x1400020b7c0, 0x1)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/chan.go:583 +0x26c fp=0x1400020b760 sp=0x1400020b6e0 pc=0x10465459c
runtime.chanrecv1(0x0?, 0x0?)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/chan.go:442 +0x14 fp=0x1400020b790 sp=0x1400020b760 pc=0x1046542f4
github.com/wailsapp/wails/v3/pkg/application.(*App).Run.func4()
        /Users/etesam/Coding/wails/v3/pkg/application/application.go:514 +0x44 fp=0x1400020b7d0 sp=0x1400020b790 pc=0x104a66bc4
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x1400020b7d0 sp=0x1400020b7d0 pc=0x1046bd114
created by github.com/wailsapp/wails/v3/pkg/application.(*App).Run in goroutine 1
        /Users/etesam/Coding/wails/v3/pkg/application/application.go:512 +0x180

goroutine 54 gp=0x14000284700 m=nil [chan receive]:
runtime.gopark(0x104ceb560, 0x1400017a718, 0xe, 0x7, 0x2)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/proc.go:402 +0xe0 fp=0x1400020bee0 sp=0x1400020beb0 pc=0x10468adb0
runtime.chanrecv(0x1400017a6c0, 0x1400020bfc0, 0x1)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/chan.go:583 +0x26c fp=0x1400020bf60 sp=0x1400020bee0 pc=0x10465459c
runtime.chanrecv1(0x0?, 0x0?)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/chan.go:442 +0x14 fp=0x1400020bf90 sp=0x1400020bf60 pc=0x1046542f4
github.com/wailsapp/wails/v3/pkg/application.(*App).Run.func5()
        /Users/etesam/Coding/wails/v3/pkg/application/application.go:520 +0x44 fp=0x1400020bfd0 sp=0x1400020bf90 pc=0x104a66ac4
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x1400020bfd0 sp=0x1400020bfd0 pc=0x1046bd114
created by github.com/wailsapp/wails/v3/pkg/application.(*App).Run in goroutine 1
        /Users/etesam/Coding/wails/v3/pkg/application/application.go:518 +0x1c0

goroutine 55 gp=0x140002848c0 m=nil [chan receive]:
runtime.gopark(0x104ceb560, 0x1400017a6b8, 0xe, 0x7, 0x2)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/proc.go:402 +0xe0 fp=0x1400020c6e0 sp=0x1400020c6b0 pc=0x10468adb0
runtime.chanrecv(0x1400017a660, 0x1400020c7c0, 0x1)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/chan.go:583 +0x26c fp=0x1400020c760 sp=0x1400020c6e0 pc=0x10465459c
runtime.chanrecv1(0x0?, 0x0?)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/chan.go:442 +0x14 fp=0x1400020c790 sp=0x1400020c760 pc=0x1046542f4
github.com/wailsapp/wails/v3/pkg/application.(*App).Run.func6()
        /Users/etesam/Coding/wails/v3/pkg/application/application.go:526 +0x44 fp=0x1400020c7d0 sp=0x1400020c790 pc=0x104a669c4
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x1400020c7d0 sp=0x1400020c7d0 pc=0x1046bd114
created by github.com/wailsapp/wails/v3/pkg/application.(*App).Run in goroutine 1
        /Users/etesam/Coding/wails/v3/pkg/application/application.go:524 +0x200

goroutine 56 gp=0x14000284a80 m=nil [chan receive]:
runtime.gopark(0x104ceb560, 0x140000345f8, 0xe, 0x7, 0x2)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/proc.go:402 +0xe0 fp=0x1400020cee0 sp=0x1400020ceb0 pc=0x10468adb0
runtime.chanrecv(0x140000345a0, 0x1400020cfb8, 0x1)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/chan.go:583 +0x26c fp=0x1400020cf60 sp=0x1400020cee0 pc=0x10465459c
runtime.chanrecv1(0x0?, 0x0?)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/chan.go:442 +0x14 fp=0x1400020cf90 sp=0x1400020cf60 pc=0x1046542f4
github.com/wailsapp/wails/v3/pkg/application.(*App).Run.func7()
        /Users/etesam/Coding/wails/v3/pkg/application/application.go:533 +0x48 fp=0x1400020cfd0 sp=0x1400020cf90 pc=0x104a668c8
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x1400020cfd0 sp=0x1400020cfd0 pc=0x1046bd114
created by github.com/wailsapp/wails/v3/pkg/application.(*App).Run in goroutine 1
        /Users/etesam/Coding/wails/v3/pkg/application/application.go:531 +0x240

goroutine 691 gp=0x14000285180 m=nil [chan receive]:
runtime.gopark(0x104ceb560, 0x140002396d8, 0xe, 0x7, 0x2)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/proc.go:402 +0xe0 fp=0x140001d1440 sp=0x140001d1410 pc=0x10468adb0
runtime.chanrecv(0x14000239680, 0x140001d1588, 0x1)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/chan.go:583 +0x26c fp=0x140001d14c0 sp=0x140001d1440 pc=0x10465459c
runtime.chanrecv2(0x140001d1598?, 0x14000289968?)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/chan.go:447 +0x14 fp=0x140001d14f0 sp=0x140001d14c0 pc=0x104654314
github.com/wailsapp/wails/v3/pkg/application.(*OpenFileDialogStruct).PromptForMultipleSelection(0x140002903f0?)
        /Users/etesam/Coding/wails/v3/pkg/application/dialogs.go:299 +0x148 fp=0x140001d15b0 sp=0x140001d14f0 pc=0x104a6d1d8
main.(*GreetService).Greet(0x140001d1608?, {0x1047088c0?, 0x140001d1608?})
        /Users/etesam/Coding/cmd-a-bug/greetservice.go:12 +0x38 fp=0x140001d15d0 sp=0x140001d15b0 pc=0x104a942f8
runtime.call32(0x14000297b60, 0x140002922a8, 0x0, 0x0, 0x0, 0x18, 0x140001d18c0)
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/asm_arm64.s:504 +0x78 fp=0x140001d1600 sp=0x140001d15d0 pc=0x1046bb1f8
runtime.reflectcall(0x104c137c0?, 0x140002830e0?, 0x8?, 0x2830e8?, 0x140?, 0x12?, 0x104c137c0?)
        <autogenerated>:1 +0x34 fp=0x140001d1640 sp=0x140001d1600 pc=0x1046c0934
reflect.Value.call({0x104c38dc0?, 0x10514ea60?, 0x213?}, {0x104aa7b8b, 0x4}, {0x14000289950, 0x1, 0x104a6b094?})
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/reflect/value.go:596 +0x670 fp=0x140001d1c90 sp=0x140001d1640 pc=0x104714df0
reflect.Value.Call({0x104c38dc0?, 0x10514ea60?, 0x10?}, {0x14000289950, 0x1, 0x1})
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/reflect/value.go:380 +0x74 fp=0x140001d1d00 sp=0x140001d1c90 pc=0x104714674
github.com/wailsapp/wails/v3/pkg/application.(*BoundMethod).Call(0x140001b4280, {0x104cf4df8, 0x140003302d0}, {0x1400032f188, 0x1, 0x14000332740?})
        /Users/etesam/Coding/wails/v3/pkg/application/bindings.go:329 +0x3e8 fp=0x140001d1ea0 sp=0x140001d1d00 pc=0x104a6b0e8
github.com/wailsapp/wails/v3/pkg/application.(*MessageProcessor).processCallMethod.func1()
        /Users/etesam/Coding/wails/v3/pkg/application/messageprocessor_call.go:107 +0xa8 fp=0x140001d1fd0 sp=0x140001d1ea0 pc=0x104a73f38
runtime.goexit({})
        /opt/homebrew/Cellar/go/1.22.2/libexec/src/runtime/asm_arm64.s:1222 +0x4 fp=0x140001d1fd0 sp=0x140001d1fd0 pc=0x1046bd114
created by github.com/wailsapp/wails/v3/pkg/application.(*MessageProcessor).processCallMethod in goroutine 690
        /Users/etesam/Coding/wails/v3/pkg/application/messageprocessor_call.go:98 +0x2e0

r0      0x0
r1      0x0
r2      0x0
r3      0x0
r4      0x180c25377
r5      0x16f822b40
r6      0x6e
r7      0x880
r8      0x121d7a31a2975ef6
r9      0x121d7a30cd156ef6
r10     0x200
r11     0xb
r12     0x0
r13     0x40500011
r14     0xffffffff
r15     0x0
r16     0x148
r17     0x1e05c32f8
r18     0x0
r19     0x6
r20     0x16f823000
r21     0x14b07
r22     0x16f8230e0
r23     0x11720a0f0
r24     0x116d474a0
r25     0x1db746888
r26     0x11720a0f0
r27     0x0
r28     0x0
r29     0x16f822ab0
lr      0x180c69cc0
sp      0x16f822a90
pc      0x180c320dc
fault   0x180c320dc
  ERROR   task: Failed to run task "run": exit status 2
```



### `wails3 doctor` output


```
etesam@Etesams-Air ~/C/cmd-a-bug (main)> wails3 doctor

                                
          Wails Doctor          
                                

                                                                                              
# System
┌──────────────────────────┐
| Name          | MacOS    |
| Version       | 14.3.1   |
| ID            | 23D60    |
| Branding      | Sonoma   |
| Platform      | darwin   |
| Architecture  | arm64    |
| Apple Silicon | true     |
| CPU           | Apple M2 |
| CPU           | Unknown  |
| GPU           | Unknown  |
| Memory        | Unknown  |
└──────────────────────────┘

# Build Environment
┌─────────────────────────────────────────────────────────┐
| Wails CLI    | v3.0.0-alpha.4                           |
| Go Version   | go1.22.2                                 |
| Revision     | 2a056a78559258d4ecb467a4f2ed48db46a44043 |
| Modified     | false                                    |
| -buildmode   | exe                                      |
| -compiler    | gc                                       |
| CGO_CFLAGS   |                                          |
| CGO_CPPFLAGS |                                          |
| CGO_CXXFLAGS |                                          |
| CGO_ENABLED  | 1                                        |
| CGO_LDFLAGS  |                                          |
| GOARCH       | arm64                                    |
| GOOS         | darwin                                   |
| vcs          | git                                      |
| vcs.modified | false                                    |
| vcs.revision | 2a056a78559258d4ecb467a4f2ed48db46a44043 |
| vcs.time     | 2024-05-25T06:33:43Z                     |
└─────────────────────────────────────────────────────────┘

# Dependencies
┌────────────────────────────────────────────────────────────────────────┐
| Xcode cli tools | 2405                                                 |
| npm             | 10.2.4                                               |
| *NSIS           | Not Installed. Install with `brew install makensis`. |
└─────────────────────── * - Optional Dependency ────────────────────────┘

# Diagnosis
 SUCCESS  Your system is ready for Wails development!

```
