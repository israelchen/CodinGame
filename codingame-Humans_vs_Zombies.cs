using System;
using System.Linq;
using System.IO;
using System.Text;
using System.Collections;
using System.Collections.Generic;

/**
 * Save humans, destroy zombies!
 **/

struct ZombieToHuman
{
    public int HumanId;
    public int ZombieId;
    public int Distance;
    public int TurnsToReach;
}

struct AshToHuman
{
    public int HumanId;
    public int Distance;
    public int TurnsToReach;
}

struct Zombie
{
    public int X;
    public int Y;
    public int NextX;
    public int NextY;
}

struct Human
{
    public int X;
    public int Y;
}

class Player
{
    private static int CalcDistance(int x1, int y1, int x2, int y2)
    {
        // d = sqrt( sqr(x2-x1) + sqr(y2 - y1) )
        return (int)Math.Sqrt(Math.Pow(x2 - x1, 2) + Math.Pow(y2 - y1, 2));
    }
    static void Main(string[] args)
    {
        string[] inputs;

        // game loop
        while (true)
        {
            var humans = new Dictionary<int, Human>();
            var zombies = new Dictionary<int, Zombie>();

            inputs = Console.ReadLine().Split(' ');
            int x = int.Parse(inputs[0]);
            int y = int.Parse(inputs[1]);

            int humanCount = int.Parse(Console.ReadLine());
            for (int i = 0; i < humanCount; i++)
            {
                inputs = Console.ReadLine().Split(' ');
                int humanId = int.Parse(inputs[0]);
                int humanX = int.Parse(inputs[1]);
                int humanY = int.Parse(inputs[2]);

                humans.Add(humanId, new Human
                {
                    X = humanX,
                    Y = humanY,
                });
            }
            int zombieCount = int.Parse(Console.ReadLine());
            for (int i = 0; i < zombieCount; i++)
            {
                inputs = Console.ReadLine().Split(' ');
                int zombieId = int.Parse(inputs[0]);
                int zombieX = int.Parse(inputs[1]);
                int zombieY = int.Parse(inputs[2]);
                int zombieXNext = int.Parse(inputs[3]);
                int zombieYNext = int.Parse(inputs[4]);

                zombies.Add(zombieId, new Zombie
                {
                    X = zombieX,
                    Y = zombieY,
                    NextX = zombieXNext,
                    NextY = zombieYNext,
                });
            }

            var zombiesToHumans = new List<ZombieToHuman>();
            var ashToHumans = new Dictionary<int, AshToHuman>();

            foreach (var human in humans)
            {
                var distance = CalcDistance(human.Value.X, human.Value.Y, x, y);

                ashToHumans.Add(human.Key, new AshToHuman
                {
                    Distance = distance,
                    HumanId = human.Key,
                    TurnsToReach = (distance - 2000) / 1000,
                });

                foreach (var zombie in zombies)
                {
                    distance = CalcDistance(zombie.Value.X, zombie.Value.Y, human.Value.X, human.Value.Y);

                    zombiesToHumans.Add(new ZombieToHuman
                    {
                        Distance = distance,
                        ZombieId = zombie.Key,
                        HumanId = human.Key,
                        TurnsToReach = (distance - 400) / 400,
                    });
                }
            }

            var h = zombiesToHumans.OrderBy(zh => zh.TurnsToReach).Where(zh => zh.TurnsToReach >= ashToHumans[zh.HumanId].TurnsToReach).First();

            // Write an action using Console.WriteLine()
            // To debug: Console.Error.WriteLine("Debug messages...");
            Console.WriteLine("{0} {1}", zombies[h.ZombieId].NextX, zombies[h.ZombieId].NextY); // Your destination coordinates
        }
    }
}
