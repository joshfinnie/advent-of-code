defmodule Computer do
  def run(a, index \\ 0) do
    if Enum.at(a, index) == 99 do
      a
    else
      op = Enum.at(a, index)
      valloc1 = Enum.at(a, index + 1)
      val1 = Enum.at(a, valloc1)
      valloc2 = Enum.at(a, index + 2)
      val2 = Enum.at(a, valloc2)
      loc = Enum.at(a, index + 3)
      a = calc(a, op, val1, val2, loc)
      run(a, index + 4)
    end
  end

  defp calc(a, op, val1, val2, loc) do
    cond do
      op == 1 ->
        val = val1 + val2
        List.replace_at(a, loc, val)
      op == 2 ->
        val = val1 * val2
        List.replace_at(a, loc, val)
    end
  end
end
