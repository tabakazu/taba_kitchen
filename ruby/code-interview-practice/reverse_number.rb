# 数値を逆順にするメソッドを作れ
# ※ 文字列にキャストは禁止

def f(n)
  r = 0
  while (n > 0)
    r = r * 10 + n % 10
    n = n / 10
  end
  r
end

p f(1234)
p f(1234567890)
p f(1111)
