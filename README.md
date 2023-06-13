# go-kafka

## folder

Go 目錄結構
/cmd
本專案的主要應用程式。

每個應用程式的目錄名應該與你的執行檔案名稱一致 (例如：/cmd/myapp)。

不要在這個目錄下放置太多程式碼。如果你認為某些程式碼也可以從其他應用程式或專案中匯入使用，那麼這些程式碼應該位於 /pkg 目錄中。如果程式碼不是可重複利用的，或者你不希望其他人使用它，請將該程式碼放到 /internal 目錄下。你未來將會驚訝的發現別人是怎麼使用你的程式碼，所以請現在就明確的表達你的意圖！

通常主要應用程式只會有一個小小的 main 函式，然後大部分的程式都是從 /internal 和 /pkg 匯入呼叫並執行，除此之外應該什麼都沒有！

請查看 /cmd 目錄獲得更多範例。

/internal
私有應用程式和函式庫的程式碼，是你不希望其他人在其應用程式或函式庫中匯入的程式碼。請注意：這個目錄結構是由 Go 編譯器本身所要求的。有關更多細節，請參閱 Go 1.4 的 release notes。注意：這個目錄並不侷限於放在專案最上層的 internal 目錄。事實上，你在專案目錄下的任何子目錄都可以包含 internal 目錄。

你可以選擇性的加入一些額外的目錄結構到你的內部套件(internal package)中，用來區分你想「共用」與「非共用」的內部程式碼(internal code)。這不是必要的（尤其是對小型專案來說），但有視覺上的線索來表達套件的共用意圖來說，肯定會更好(nice to have)。你的應用程式程式碼可以放在 /internal/app 目錄下 (例如：/internal/app/myapp)，而這些應用程式共享的程式碼就可以放在 /internal/pkg 目錄下 (例如：/internal/pkg/myprivlib)。

/pkg
函式庫的程式碼當然可以讓外部應用程式來使用 (例如：/pkg/mypubliclib)，其他專案會匯入這些函式庫，並且期待它們能正常運作，所以要把程式放在這個目錄下請多想個幾遍！:-) 注意：使用 internal 目錄可以確保私有套件不會被匯入到其他專案使用，因為它是由 Go 的編譯器強制執行的，所以是比較好的解決方案。使用 /pkg 目錄仍然是一種很好的方式，它代表其他專案可以安全地使用這個目錄下的程式碼。由 Travis Jeffery 撰寫的 I'll take pkg over internal 文章提供了關於 pkg 和 internal 目錄很好的概述，以及使用它們的時機點。

當專案的根目錄包含許多不是用 Go 所寫的元件與目錄時，將 Go 程式碼放在一個集中的目錄下也是種不錯的方法，這使得運行各種 Go 工具變得更加容易（正如以下這些演講中提到的那樣：來自 GopherCon EU 2018 的 Best Practices for Industrial Programming、GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps 和 GoLab 2018 - Massimiliano Pippi - Project layout patterns in Go）。

如果你想查看哪些知名的 Go 專案使用本專案的目錄結構，請查看 /pkg 目錄。這是一組常見的目錄結構，但並不是所有人都接受它，有些 Go 社群的人也不推薦使用。

如果你的應用程式專案真的很小，或是套用這些資料夾不會對你有太大幫助（除非你真的很想用XD），不使用本專案推薦的目錄結構是完全沒問題的。當你的專案變的越來越大，根目錄將會會變得越來越複雜（尤其是當你有許多不是 Go 所寫的元件時)，你可以考慮參考這個專案所建議的目錄結構來組織你的程式碼。

/vendor
應用程式的相依套件可透過手動管理，或使用你喜歡的相依性套件管理工具，例如內建的 Go Modules 特性。使用 go mod vendor 命令可以幫你建立一個 /vendor 目錄。請注意：如果你不是用 Go 1.14+ 版本的話，你可能需要在執行 go build 的時候增加 -mod=vendor 命令列參數。從 Go 1.14 開始，這個參數預設就是啟用的。

如果你正在建立一個函式庫套件，那麼請不要將你應用程式的相依套件加入版控！

注意：從 Go 1.13 開始，Go 預設啟用了模組的代理伺服器 (module proxy) 功能 (預設使用 https://proxy.golang.org 作為模組的代理伺服器)。你可以從這裡查看這功能是否符合你的需求與限制。如果你可以使用 module proxy 的話，那麼你根本不需要 vendor 目錄。

服務應用程式目錄 (Service Application Directories)
/api
OpenAPI/Swagger 規格、JSON schema 檔案、各種協定定義檔。

請查看 /api 目錄獲得更多範例。

Web 應用程式目錄 (Web Application Directories)
/web
Web 應用程式相關的元件：靜態 Web 檔案、伺服器端範本與 SPAs 相關檔案。

通用應用程式目錄 (Common Application Directories)
/configs
組態設定的檔案範本或預設設定。

將你的 confd 或 consul-template 範本檔案放在這裡。

/init
放置 System init (systemd, upstart, sysv) 與 Process manager/supervisor (runit, supervisor) 相關設定。

/scripts
放置要執行各種建置、安裝、分析等操作的命令腳本！

這些腳本可以讓你在根目錄的 Makefile 更小、更簡單（例如：https://github.com/hashicorp/terraform/blob/master/Makefile)。

請查看 /scripts 目錄獲得更多範例。

/build
封裝套件與持續整合(CI)。

將你的雲端 (AMI)、容器 (Docker)、OS (deb, rpm, pkg) 套件的組態設定與腳本放在 /build/package 目錄下。

將你的 CI (Travis CI, CircleCI, Drone CI) 的組態設定與腳本放在 /build/ci 目錄中。請注意：有些 CI 工具 (例如 Travis CI 等)，它們對這些組態設定檔案的位置非常挑剔。如果可能的話，請嘗試將檔案放在 /build/ci 目錄中，並連結 (linking) 這些檔案到 CI 工具期望它們出現的位置。

/deployments
IaaS、PaaS、系統和容器編配部署的組態設定與範本 (docker-compose、kubernetes/helm、mesos、terraform、bosh)。注意：在某些儲存庫中（特別是那些部署在 kubernetes 的應用程式），這個目錄會被命名為 /deploy。

/test
額外的外部測試應用程式和測試資料。你可以自在的調整你在 /test 目錄中的結構。對於較大的專案來說，通常會有一個 data 資料夾也是蠻正常的。例如：如果你需要 Go 忽略這些目錄下的檔案，你可以使用 /test/data 或 /test/testdata 當作你的目錄名稱。請注意：Go 還會忽略以 . 或 _ 開頭的目錄或檔案，所以你在測試資料的目錄命名上，將擁有更大的彈性。

請查看 /test 目錄獲得更多範例。

其他目錄
/docs
設計和使用者文件 (除了 godoc 自動產生的文件之外)。

請查看 /docs 目錄獲得更多範例。

/tools
這個專案的支援工具。注意：這些工具可以從 /pkg 和 /internal 目錄匯入程式碼。

請查看 /tools 目錄獲得更多範例。

/examples
放置關於你的應用程式或公用函式庫的使用範例。

請查看 /examples 目錄獲得更多範例。

/third_party
外部輔助工具、Forked 程式碼，以及其他第三方工具 (例如：Swagger UI)。

/githooks
Git hooks。

/assets
其他要一併放入儲存庫的相關檔案 (例如圖片、Logo、... 等等)。

/website
如果你不使用 GitHub Pages 的話，這裡可以放置專案的網站相關資料。

請查看 /website 目錄獲得更多範例。

## reference

- https://www.baeldung.com/ops/kafka-docker-setup
- https://github.com/golang-standards/project-layout/blob/master/README_zh-TW.md