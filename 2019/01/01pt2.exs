defmodule Fuel do
  def calc(total, weight) do
    weight = div(weight, 3) - 2
    if weight <= 0 do
      total
    else
      total = total + weight
      calc(total, weight)
    end
  end
end

IO.puts File.read!(Path.join(__DIR__, "input.txt"))
    |> String.split("\n", trim: true)
    |> Enum.map(&String.to_integer/1)
    |> Enum.map(&(Fuel.calc(0, &1)))
    |> Enum.sum
