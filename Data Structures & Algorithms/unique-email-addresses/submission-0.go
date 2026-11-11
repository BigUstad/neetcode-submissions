func numUniqueEmails(emails []string) int {
    m := make(map[string]bool)
    for _, e := range emails {
        eSlc := strings.Split(e, "@")
        eSlc[0] = strings.ReplaceAll(eSlc[0], ".", "")
        eSlc2 := strings.Split(eSlc[0], "+")
        mEle := eSlc2[0] + "@" + eSlc[1]
        //fmt.Println(mEle)
        m[mEle] = true
    }
    return len(m)
}