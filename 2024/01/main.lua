local path = "/home/simon/Projects/Fun/advent-of-code/2024/01/input.txt"
local f = io.open(path, "r")

local left = {}
local right = {}

-- Create two sets for left and right to prevent duplicates.
--      Time Complexity:     Ο(n) Ω(n) Θ(n)
--      Space Complexity:    Ο(n) Ω(n) Θ(n)

for line in io.lines(path) do
	for num in string.gmatch(line, "%d+") do
	end
end

-- Sort the sets to make it easy to compute.
-- Create two sets for left and right to prevent duplicates.
--      Time Complexity:     Ο(n) Ω(n) Θ(n)
--      Space Complexity:    Ο(n) Ω(n) Θ(n)

-- Iterate through the values and calculate the sum.
--      Time Complexity:     Ο(n) Ω(n) Θ(n)
--      Space Complexity:    Ο(n) Ω(n) Θ(n)

io.close(f)
