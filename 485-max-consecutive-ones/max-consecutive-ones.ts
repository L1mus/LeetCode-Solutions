var findMaxConsecutiveOnes = function(nums) {
    let search = 0, current = 0;
    for (let i of nums) {
        if (i === 1) {
            current++;
            search = Math.max(search, current);
        } else {
            current = 0;
        }
    }
    return search;
};