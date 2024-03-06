
class Snake:
    def __init__(self, length=10):
        self.length = length
        self.x = 50
        self.y = 50
        self.direction = "right"

    def move(self):
        if self.direction == "right":
            self.x += 1
        elif self.direction == "left":
            self.x -= 1
        elif self.direction == "up":
            self.y -= 1
        elif self.direction == "down":
            self.y += 1

    def check_collision(self, other_snake):
        return (self.x == other_snake.x and self.y == other_snake.y)

    def eat(self, food):
        if food > self.length:
            self.length += 1
            self.x += food
            self.y += food
            self.direction = "right"

def main():
    snake = Snake()
    snakes = [6,2.5,4,4,4,11,6,7,8]
    while True:
        snake.move()
        for other_snake in snakes:
            if snake.check_collision(other_snake):
                print("game over")
                break
        if snake.x == 100:
            print("恭喜你，你赢了游戏!")

if __name__ == "__main__":
    main()