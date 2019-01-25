defmodule Solution do

  def array(n) when n > 0 do
    array(n) ++ [n]
  end

  def array(_n) do
    []
  end

  def main do
    n = IO.read(:stdio, :line) |> String.trim() |> String.to_integer
    array(n) |> IO.inspect(charlists: false)
  end
end

Solution.main
