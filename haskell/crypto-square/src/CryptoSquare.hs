module CryptoSquare (encode, findRectangleC, splitAll) where

import Data.Char ( isAlpha, toLower )

encode :: String -> String
encode xs =
    let normalized = normalize xs
        rectangleC = findRectangleC $ length normalized
        rectangle = splitAll rectangleC normalized
    in normalized

normalize :: String -> String
normalize xs = filter isAlpha $ map toLower xs

findRectangleC :: Int -> Int
findRectangleC len = ceiling $ sqrt $ fromIntegral len

splitAll :: Int -> String -> [String]
splitAll _ "" = []
splitAll n xs =
    let (f, s) = splitAt n xs
    in f : splitAll n s
