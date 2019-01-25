defmodule Solution do
  # Enter your code here. Read input from STDIN. Print output to STDOUT
  def print(times, elem) when times > 0 do
    IO.puts elem
    print(times-1, elem)
  end

  def print(_times, _elem) do
    # Do nothing
  end

  def read_lines(n) do
    elem = IO.read(:stdio, :line) |> String.trim
    case elem do
      "" -> ""
      val ->
        print(n, val)
        read_lines(n)
    end
  end

  def main do
    times = IO.read(:stdio, :line) |> String.trim |> String.to_integer
    read_lines(times)
  end
end
