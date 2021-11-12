module Anagram (anagramsFor) where

import Data.List ( sort )
import Data.Char (toLower)

anagramsFor :: String -> [String] -> [String]
anagramsFor word xxs = foldr ((++) . anagramFor word) [] xxs

anagramFor :: String -> String -> [String]
anagramFor word cand =
    [cand | sort lowerWord == sort lowerCand && lowerWord /= lowerCand ]
    where lowerWord = map toLower word
          lowerCand = map toLower cand