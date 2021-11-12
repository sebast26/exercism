module School (School, add, empty, grade, sorted) where

import Data.Maybe ( fromMaybe )
import Data.List as List ( sort )
import qualified Data.Map as Map

type School = Map.Map Int [String]

add :: Int -> String -> School -> School
add gradeNum student = Map.insertWith (\new old -> sort (new ++ old)) gradeNum [student]

empty :: School
empty = Map.empty

grade :: Int -> School -> [String]
grade gradeNum school =
    fromMaybe [] (Map.lookup gradeNum school)

sorted :: School -> [(Int, [String])]
sorted = Map.toList
