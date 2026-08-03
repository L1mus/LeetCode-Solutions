function theMaximumAchievableX(num: number, t: number): number {
    let achievableNumX : number = 0;
    for(let i: number = 0 ; i < t ; i++){
        num++
        achievableNumX--
    }
    return -(achievableNumX)+num;
};