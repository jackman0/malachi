# Malachi

Malachi, named for the only bowling song I know, is a simple tool to track your
bowling score.

I created this as a coding exercise, and you should consider it a first draft. I
learned how bowling scores actually worked as I coded it. Consequently, I was 
pretty fast and loose with validation and user input. Nevertheless, I think it's
a good place to start for discussion.

## Usage

The application takes data via STDIN, and displays the current frame and score
on STDOUT.

Frame scores are entered in comma-separated format. For a strike, you may enter
only 'X' or '10'. Spares may be represented numerically or with the spare symbol
(slash), such as '2,8' or '2,/'. The final frame will accept a triple, such as
'2,/,3' or 'X,2,3'.

The score input either works or it doesn't -- validation does not generate any
meaningful output.

## TODO

 * Display the entire score sheet
 * Support multiple players with names
 * Provide more useful input validation response
 * Make validation more strict, especially when it comes to the final frame.
 * Make editing possible.

I really wouldn't mind turning this into a web app or something with a discrete
front-end and back-end. I can easily imagine this being turned into a tool to 
permanently store scorecards for historical reference, and a way for people to
play a game with each other remotely.