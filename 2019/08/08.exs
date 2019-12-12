defmodule Part08 do
  def part1(input, {w, h}) do
    IO.puts least_corrupted_layer(input, {w, h})
  end

  def part2(input, {w, h}) do
    IO.puts render_pixels(input, {w, h})
  end

  defp least_corrupted_layer(input, {w, h}) do
    input
      |> Enum.chunk_every(w * h)
      |> Enum.min_by(&count(&1, 0))
      |> get_ones_and_twos()
  end

  defp render_pixels(input, {w, h}) do
    input
      |> Enum.chunk_every(w* h)
      |> List.zip()
      |> Enum.map(&Tuple.to_list/1)
      |> Enum.map(&render_pixel/1)
      |> Enum.chunk_every(w)
      |> Enum.map(&Enum.join/1)
      |> Enum.map(&String.trim_trailing/1)
      |> Enum.join("\n")
  end

  defp render_pixel(pixel_layers) do
    pixel_layers
    |> Enum.map(&replace_black_with_space/1)
    |> Enum.find(&(&1 != 2))
  end

  defp replace_black_with_space(0), do: " "
  defp replace_black_with_space(other), do: other

  defp count(layer, value) do
    layer
      |> Enum.count(&(&1==value))
  end

  def get_ones_and_twos(layer) do
    count(layer, 1) * count(layer, 2)
  end
end

input = File.read!(Path.join(__DIR__, "input.txt"))
  |> String.graphemes()
  |> Enum.map(&String.to_integer/1)

IO.puts Part08.part1(input, {25, 6})
IO.puts Part08.part2(input, {25, 6})
