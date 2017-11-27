# snvalidator
A validator for Chinese id card number.

### Usage

```golang
import "github.com/songjiayang/snvalidator"

func main(){
  sn := "513826199104253418"
  valid, _ := validator.IsValidSN(sn) // valid should be true
}
```

### References

[中华人民共和国公民身份号码](https://zh.wikipedia.org/wiki/中华人民共和国公民身份号码)
