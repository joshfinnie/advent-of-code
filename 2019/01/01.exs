IO.puts File.read!(Path.join(__DIR__, "input.txt"))
    |> String.split("\n", trim: true)
    |> Enum.map(&String.to_integer/1)
    |> Enum.map(&(div(&1, 3) - 2))
    |> Enum.sum
