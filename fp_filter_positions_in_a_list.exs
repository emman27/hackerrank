defmodule Solution do
  def remove_odd([hd | tl], include) do
    case include do
      true -> [hd] ++ remove_odd(tl, not include)
      false -> remove_odd(tl, not include)
    end
  end

  def remove_odd(_lst, _include) do
    []
  end

  def read_input() do
    case IO.read(:stdio, :line) do
      :eof -> []
      line -> [line |> String.trim] ++ read_input()
    end
  end
end

Solution.read_input() |> Solution.remove_odd(false) |> Enum.map(fn val -> IO.puts(val) end)
