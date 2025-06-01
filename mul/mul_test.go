package mul
import "testing"
func TestAdd(t *testing.T){
	got := Mul(2,3)
	want:=6
	if got != want {
		
        t.Errorf("Mul(2, 3) = %d; want %d", got, want)
        
	}
}