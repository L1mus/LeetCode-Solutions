function scoreOfString(s: string): number {
    let n = s.length
    let sum = 0

    for (let i = 1; i < n; i++) {
        let p = s[i - 1]
        let c = s[i];

        sum += Math.abs(p.charCodeAt(0) - c.charCodeAt(0))
    }

    return sum
};