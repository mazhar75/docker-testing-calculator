package mod
import "testing"
func TestAdd(t *testing.T){
	got,err := Mod(18,3)
	want:=0
	if err!=nil && got != want {
		
        t.Errorf("Add(2, 3) = %d; want %d", got, want)
        
	}
}