a = File.read!(Path.join(__DIR__, "input.txt"))
    |> String.split([",", "\n"], trim: true)
    |> Enum.map(&String.to_integer/1)

IO.puts Enum.at(Computer.run(a), 0)
