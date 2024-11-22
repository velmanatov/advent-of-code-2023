# day 1 wants us to find the first and last digits in every line (bearing in mind they may be the same one), concatenate them to make a two digit number and sum them all

file = File.open("input.txt")
lines = file.readlines.map(&:chomp)

regex = /\d/
sum = 0

lines.each do |line|
    # scan for all numerical digits in the string using the regex. Then just concatenaate them together, remembering to multiple the first by 10
    allDigits = line.scan(regex)
    sum += Integer(allDigits.first) * 10 + Integer(allDigits.last)
end

puts(sum)