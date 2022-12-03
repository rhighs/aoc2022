defmodule Day3 do
  def read_file do
    {:ok, content} = File.read("input.txt")
    content
  end

  def score(ascii_value) do
    cond do
      ascii_value > 96 && ascii_value < 123 ->
        ascii_value - 96
      ascii_value > 64 && ascii_value < 91 ->
        27 + ascii_value - 65
    end
  end

  def p1(lines) do
    lines
      |> Enum.map(fn(sack) -> String.split_at(sack, div(String.length(sack), 2)) |> Tuple.to_list |> Enum.map(&(MapSet.new((String.graphemes(&1))))) end)
      |> Enum.map(fn([c1, c2]) -> MapSet.intersection(c1, c2) |> MapSet.to_list |> Enum.map(&(score(:binary.first &1))) end) |> List.flatten
      |> Enum.sum
  end

  def p2(lines) do
    lines
    |> Stream.chunk_every(3, 3, :discard) |> Enum.map(fn(group) -> Enum.map(group, &(MapSet.new((String.graphemes(&1))))) end)
    |> Enum.map(fn([e1, e2, e3]) -> MapSet.intersection(e1, e2) |> MapSet.intersection(e3) |> MapSet.to_list |> Enum.map(&(score(:binary.first &1))) end) |> List.flatten
    |> Enum.sum
  end
end

lines = Day3.read_file
  |> String.split("\n")

IO.puts "p1: #{Day3.p1(lines)}"
IO.puts "p2: #{Day3.p2(lines)}"
