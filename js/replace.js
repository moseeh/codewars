function removeExclamationMarks(s) {
    let newstr = '';
    for(let i = 0; i < s.length; i++) {
        if(s[i] !== '!') {
            newstr += s[i];
        }
    }
    return newstr;
}