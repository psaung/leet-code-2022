package main


func generateParenthesis(n int) []string {
    result := []string{}
    
    if n == 0 {
        return result
    }
    
    var backtrack func (string, int, int, int)
    
    backtrack = func(s string, open int, closed int, n int) {
        
        if open == n && closed == n {
            result = append(result, s)
        }
        
        if open < n {
            backtrack(s + "(" , open + 1, closed, n)
        }
        
        if closed < open {
            backtrack(s + ")" , open, closed + 1, n)
        }
    }
    
    backtrack("", 0, 0, n)
    
    return result;
    
}