defmodule Solution do

  def read_input do
    case IO.read(:stdio, :line) do
      :eof -> []
      val ->
        [String.trim(val) |> String.to_integer] ++ read_input()
    end
  end
end

delimiter = IO.read(:stdio, :line) |> String.trim |> String.to_integer
Solution.read_input() |> Enum.filter(fn val -> val < delimiter end) |> Enum.map(fn val -> IO.puts(val) end)
