function getSneakyNumbers(nums: number[]): number[] {
    const seen = new Set<number>();
    const duplicates = new Set<number>();

    for (const num of nums) {
        if (seen.has(num)) {
            duplicates.add(num);
        } else {
            seen.add(num);
        }
    }
    return [...duplicates]
};