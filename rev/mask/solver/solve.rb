str_a = "atd4`qdedtUpetepqeUdaaeUeaqau"
str_b = "c`b bk`kj`KbababcaKbacaKiacki"

flag = str_a.chars.zip(str_b.chars).map do |(a, b)|
  (a.ord | b.ord).chr
end.join

puts flag