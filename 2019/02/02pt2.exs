defmodule Manipulate do
  def array(a, i, j) do
    a = List.replace_at(a, 1, i)
    a = List.replace_at(a, 2, j)
    val = Enum.at(Computer.run(a), 0)
    if val == 19690720 do
      IO.puts "noun: #{i} | verb: #{j}"
    end
  end
end

a = File.read!(Path.join(__DIR__, "input.txt"))
    |> String.split([",", "\n"], trim: true)
    |> Enum.map(&String.to_integer/1)

for i <- 0..99, j <- 0..99, do: Manipulate.array(a, i, j)
