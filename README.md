# snvalidator
A validator of Chinese ID card number.

### Usage

```golang
import "github.com/songjiayang/snvalidator"

func main(){
  sn := "513826199104253418"
  isValid, _ := snvalidator.IsValidSN(sn) // isValid should be true
}
```

### References

[中华人民共和国公民身份号码](https://zh.wikipedia.org/wiki/中华人民共和国公民身份号码)
