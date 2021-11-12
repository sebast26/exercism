module Clock (addDelta, fromHourMin, toString) where

data Clock = Clock Int Int deriving (Eq)

fromHourMin :: Int -> Int -> Clock
fromHourMin hour minute = Clock hours minutes
  where hours = (hour + addedHours) `mod` 24
        minutes = minute `mod` 60
        addedHours = minute `div` 60

toString :: Clock -> String
toString (Clock hour minute) = fullhour ++ ":" ++ fullminute
  where fullhour = if hour < 10 then "0" ++ show hour else show hour
        fullminute = if minute < 10 then "0" ++ show minute else show minute

addDelta :: Int -> Int -> Clock -> Clock
addDelta hour minute (Clock cHour cMinute) = Clock newHour newMinute
  where newHour = (hour + cHour + addedHour) `mod` 24
        newMinute = (minute + cMinute) `mod` 60
        addedHour = (minute + cMinute) `div` 60
