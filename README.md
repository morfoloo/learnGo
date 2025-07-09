### Aurora usage

In my projects timer.go and converter.go, I used Aurora to add colored text in the terminal.
If you want to use it too, just follow these steps:

**1**. Navigate to your project directory (mine is called `learnGo`):
```
cd learnGo/
```
**2**. Initialize a Go module:
```
go mod init learnGo
```
**3**. Install Aurora:
```
go get -u github.com/logrusorgru/aurora/v4
```

Done! You can now use
```
fmt "github.com/logrusorgru/aurora/v4"
```
