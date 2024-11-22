# day 1 wants us to find the first and last digits in every line (bearing in mind they may be the same one), concatenate them to make a two digit number and sum them all
# part two adds a twist that some digits are spelled out - like "one", "two", "three", etc.

file = File.open("input.txt")
lines = file.readlines.map(&:chomp)

numericalRegex = /\d/
# todo - I know this fails for input six3one4sixeighttwosbkqdjfhfroneights - since the regexp match does not include "overlapping" results. Need look ahead?
alphaRegex = /one|two|three|four|five|six|seven|eight|nine|ten/
numberHash = { "one" => 1, "two" => 2, "three" => 3, "four" => 4, "five" => 5, "six" => 6, "seven" => 7, "eight" => 8, "nine" => 9 }
sum = 0

FIXNUM_MAX = (2**(0.size * 8 -2) -1)

lines.each do |line|
    # scan for all numerical digits in the string using both regular expressions. Then make a two digit numer out of them, remembering to multiply the first by 10
    # thanks to https://brettterpstra.com/2023/12/17/ruby-regexp-scan-with-matchdata/ for idea on how to scan and get a set of MatchData groups
    allDigitMatches = line.to_enum(:scan, numericalRegex).map { Regexp.last_match }
    allSpelledOutDigitMatches = line.to_enum(:scan, alphaRegex).map { Regexp.last_match }

    firstNumberPosition = allDigitMatches.length() > 0 ? allDigitMatches[0].begin(0) : FIXNUM_MAX
    firstSpelledOutPosition = allSpelledOutDigitMatches.length() > 0 ? allSpelledOutDigitMatches[0].begin(0) : FIXNUM_MAX

    if (firstNumberPosition < firstSpelledOutPosition)
        firstNumber = Integer(allDigitMatches.first[0])
    else
        firstNumber = numberHash[allSpelledOutDigitMatches.first[0]]
    end

    lastNumberPosition = allDigitMatches.length() > 0 ? allDigitMatches.last.begin(0) : -1
    lastSpelledOutPosition = allSpelledOutDigitMatches.length() > 0 ? allSpelledOutDigitMatches.last.begin(0) : -1

    if (lastNumberPosition > lastSpelledOutPosition)
        lastNumber = Integer(allDigitMatches.last[0])
    else
        lastNumber = numberHash[allSpelledOutDigitMatches.last[0]]
    end    

    sum += firstNumber * 10 + lastNumber
end

puts(sum)
