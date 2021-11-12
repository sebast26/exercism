module Phone (number) where

import Data.Char ( isAlpha )

number :: String -> Maybe String
number = filter isAlpha