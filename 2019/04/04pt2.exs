defmodule Password do
  def run(input) do
    input
        |> String.split(["-", "\n"], trim: true)
        |> Enum.map(&String.to_integer/1)
        |> (fn [a, b] -> a..b end).()
        |> Enum.to_list()
        |> Enum.filter(&test/1)
        |> Enum.count()
  end

  def test(num) do
    Integer.digits(num)
      |> (&(increases(&1) && has_separate_doubles(&1))).()
  end

  defp increases([_x]), do: true
  defp increases([x, y | _rest]) when x > y, do: false
  defp increases([_x, y | rest]), do: increases([y | rest])

  defp has_separate_doubles([]), do: false
  defp has_separate_doubles([_x]), do: false
  defp has_separate_doubles([x, x, x, x, x | _rest]), do: false
  defp has_separate_doubles([x, x, x, x | rest]), do: has_separate_doubles(rest)
  defp has_separate_doubles([x, x, x | rest]), do: has_separate_doubles(rest)
  defp has_separate_doubles([x, x | _rest]), do: true
  defp has_separate_doubles([_x, _y]), do: false
  defp has_separate_doubles([_x, y | rest]), do: has_separate_doubles([y | rest])

end

input = IO.gets "input? " |> String.trim
IO.puts Password.run(input)
