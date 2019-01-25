defmodule Solution do
    a = IO.read(:stdio, :line) |> String.trim |> String.to_integer
    b = IO.read(:stdio, :line) |> String.trim |> String.to_integer
    IO.puts a + b
end
