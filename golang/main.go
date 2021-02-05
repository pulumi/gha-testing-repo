package main

import (
	random "github.com/pulumi/pulumi-random/sdk/v2/go/random"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := random.NewRandomPet(ctx, "my-user-name", &random.RandomPetArgs{})
		if err != nil {
			return err
		}

		_, err = random.NewRandomPassword(ctx, "my-password", &random.RandomPasswordArgs{
			Length: pulumi.Int(30),
		})

		return nil
	})
}
