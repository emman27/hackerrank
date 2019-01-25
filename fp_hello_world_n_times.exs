defmodule Solution do
  def hello_world(n) when n > 0 do
    IO.puts "Hello World"
    hello_world(n-1)
  end

  def hello_world(_n) do
  end

  def main do
    IO.read(:stdio, :line) |> String.trim |> String.to_integer |> hello_world
  end
end
