package div
import "testing"
func TestAdd(t *testing.T){
	got,err := Div(2,0)
	want:=-1
	if err==nil {
		
        t.Errorf("Add(2, 3) = %d; want %d", got, want)
        
	}
}