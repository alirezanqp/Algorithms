function binarySearch<T>(list: T[], item: T): number | null {
    let low: number = 0;
    let high: number = list.length - 1;

    while (low <= high) {
        let mid: number = Math.floor((low + high) / 2);
        let guess: T = list[mid]

        if (guess === item) {
            return mid;
        }

        if (guess < item) {
            low = mid + 1;
        } else {
            high = mid - 1;
        }
    }

    return null;
}

const myList = [1, 3, 5, 6, 9];

console.log(binarySearch(myList, 6));
console.log(binarySearch(myList, 8));
