module Acronym (abbreviate) where

import Data.Char ( isUpper, toUpper, isAlpha )

abbreviate' :: Char -> String -> String 
abbreviate' _ [] = []
abbreviate' ' ' (x:xs)
    | isAlpha x = toUpper x:abbreviate' x xs
    | otherwise = abbreviate' x xs
abbreviate' '-' (x:xs)
    | isAlpha x = toUpper x:abbreviate' x xs
    | otherwise = abbreviate' x xs
abbreviate' c (x:xs)
    | isUpper c = abbreviate' x xs
    | isUpper x = x:abbreviate' x xs
    | otherwise = abbreviate' x xs

abbreviate :: String -> String
abbreviate = abbreviate' ' '
