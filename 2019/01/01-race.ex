IO.puts IO.read(:stdio, :all)
    |> String.split("\n", trim: true)
    |> Stream.map(&String.to_integer/1)
    |> Stream.map(&(div(&1, 3) - 2))
    |> Enum.sum
