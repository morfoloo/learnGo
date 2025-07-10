### Fyne usage

In my project rtx5090 I used Fyne to add GUI. If you want to use it too, just follow these steps:

**1**. Navigate to your project directory (mine is called `learnGo`):

```
cd learnGo/
```

**2**. Initialize a Go module:
```
go mod init learnGo
```

**3**. Install fyne:

```
go get fyne.io/fyne/v2
```

***Done***! You can now use
```
import "fyne.io/fyne/v2"
```

*Also you can just download "rtx5090"(this is for linux) and run it with command*
```
./rtx5090
```

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

***Done!*** You can now use
```
import "github.com/logrusorgru/aurora/v4"
```
