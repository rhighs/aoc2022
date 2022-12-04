defmodule Day4 do
  def read_file(path) do
    {:ok, content} = File.read(path)
    content
  end

  def includes([[x1, x2], [x3, x4]]), do: x1 <= x3 && x2 >= x4 || x3 <= x1 && x4 >= x2
  def overlaps([[x1, x2], [x3, x4]]), do: (x1 <= x3 && x2 <= x4 && x2 >= x3) || (x3 <= x1 && x4 <= x2 && x4 >= x1)

  def parse_input(input) do
    input
    |> String.split("\n")
    |> Enum.map(fn(line) -> String.split(line, ",")
    |> Enum.map(fn(pair_item) -> 
      String.split(pair_item, "-")
      |> Enum.map(fn(str_value) -> 
        {int_value, ""} = Integer.parse(str_value)
        int_value
        end)
      end) 
    end)
  end

  def p1(input), do: parse_input(input) |> Enum.count(&(includes(&1)))
  def p2(input), do: parse_input(input) |> Enum.count(&(includes(&1) || overlaps(&1)))
end

IO.puts(Day4.read_file("input.txt") |> String.trim("\n") |> Day4.p1)
IO.puts(Day4.read_file("input.txt") |> String.trim("\n") |> Day4.p2)
