// represents info about a potential "part" in the input grid - i.e. a number and its position
export interface NumberMatch {
    number: number,
    x: number,
    y: number,
    length: number
}