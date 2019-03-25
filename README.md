# library-e2e-tester

A CLI testing tool that sends events to Segment and verifies events are received by a webhook connected to Segment.

## Usage

The tester is used as follows:

```bash
./tester -segment-write-key='...' -webhook-auth-username='...' -webhook-bucket='...' -path='...'
```

The tester will invoke the library CLI with some fixtures, and verify that events appeared in a connected webhook.

For a library to be tested by this tester, the library must provide a CLI interface that conforms to the contract enforced by the tester:

```bash
analytics --type=<type>
          --writeKey=<writeKey>
          --userId=<userId>
          [--event=<event> --properties=<properties>] # Track
          [--name=<name> --properties=<properties>] # Page/Screen
          [--traits=<traits>] # Identify
          [--groupId=<groupId> --traits=<traits>] # Group
```

The setup must be done manually:

1. Create a Segment workspace and project. Note the writeKey of the project.
2. Create a Webhook bucket. Note the Webhook bucket ID and Webhook auth username.
3. Add the Webhook bucket as a Segment webhook destination.

![tester](https://cldup.com/luiNQxqYu9.png)

The Webhook API is documented on [SwaggerHub](https://app.swaggerhub.com/apis/Segment.io/Webhook-E2E/1.0.1).
