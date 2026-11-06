g nums div = [x | (i, x) <- zip [1..] nums, i `mod` div > 0]
f nums@(x:y:xs) mode
    | mode == 1 = f (g nums 2) 2
    | mode == 2 = f (g (y:xs) 2) 3
    | mode == 3 = f (g nums 3) 1
f (x:xs) mode = x

main = do
    input <- readLn :: IO Int
    print $ f [1..input] 1

