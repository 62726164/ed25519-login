# Server Notes

## Signature Time limit
```
    func toTimeRange(epoch int64) []string {
        times := make([]string, 0)
 
        for i := epoch; i < epoch+30; i++ {
            s := strconv.FormatInt(i, 10)
            times = append(times, s)
        }
 
        for i := epoch; i > epoch-30; i-- {
            s := strconv.FormatInt(i, 10)
            times = append(times, s)
        }
 
        return times
    }

    ...

    now := time.Now()
    secs := now.Unix()
    times := toTimeRange(secs)

    for _, time := range times {
        verified := ed25519.Verify(dpubkey, []byte(time), dsignature)
        if verified {
            session.Values["authenticated"] = true
            session.Save(r, w)
            usedSigs[signature] = true
            return
        }
    }
```

## Prevention of replay attacks
```
     var usedSigs map[string]bool = make(map[string]bool, 0)

	 ...

     _, found := usedSigs[signature]
     if found {
         log.Printf("The signature was previously used to login.\n")
         http.Redirect(w, r, "login-failed", http.StatusSeeOther)
         return
     }
```
