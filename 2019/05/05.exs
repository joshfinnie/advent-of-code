defmodule Computer do
  def run(a, index \\ 0) do
    if Enum.at(a, index) == 99 do
      a
    else
      op = Enum.at(a, index)
      {a, i} = calc(a, op, index)
      try do
        run(a, index + i)
      rescue
        ArgumentError -> nil
      end
    end
  end

  defp calc(a, op, index) do
    cond do
      op == 1 ->
        {opt1(a, index), 4}
      op == 2 ->
        {opt2(a, index), 4}
      op == 3 ->
        {opt3(a, index), 1}
      paramTest?(op) == true ->
        {optParam(a, index), 1}
    end
  end

  defp opt1(a, index) do
      valloc1 = Enum.at(a, index + 1)
      val1 = Enum.at(a, valloc1)
      valloc2 = Enum.at(a, index + 2)
      val2 = Enum.at(a, valloc2)
      loc = Enum.at(a, index + 3)
      val = val1 + val2
      List.replace_at(a, loc, val)
  end

  defp opt2(a, index) do
      valloc1 = Enum.at(a, index + 1)
      val1 = Enum.at(a, valloc1)
      valloc2 = Enum.at(a, index + 2)
      val2 = Enum.at(a, valloc2)
      loc = Enum.at(a, index + 3)
      val = val1 * val2
      List.replace_at(a, loc, val)
  end

  defp opt3(a, index) do
      valloc1 = Enum.at(a, index + 1)
      val1 = Enum.at(a, valloc1)
      List.replace_at(a, val1, val1)
  end

  defp paramTest?(op) do
    String.match?(Integer.to_string(op), ~r/^[[:alnum:]]+$/) && (length(Integer.digits(op)) == 4 || length(Integer.digits(op)))
  end

  defp optParam(a, index) do
    IO.puts index
    a
  end

end

defmodule Part1 do
  def main do
    a = "1002,4,3,4,33"
    #a = File.read!(Path.join(__DIR__, "input.txt"))
        |> String.split([",", "\n"], trim: true)
        |> Enum.map(&String.to_integer/1)

    #IO.puts Enum.at(Computer2.run(a), 0)
    Computer.run(a)
  end
end

Part1.main()
