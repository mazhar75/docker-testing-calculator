package sub
import "testing"
func TestAdd(t *testing.T){
	got := Sub(2,3)
	want:=-1
	if got != want {
		
        t.Errorf("Sub(2, 3) = %d; want %d", got, want)
        
	}
}