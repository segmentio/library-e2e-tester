# library-e2e-tester

A CLI testing tool that sends events to Segment and verifies events are received by a webhook connected to Segment.

# Usage

The tester is used as follows:

```
./tester -segment-write-key='...' -runscope-token='...' -webhook-bucket='...' -path='...''
```

The tester will invoke the library CLI with some fixtures, and verify that events appeared in a Runscope connected webhook.

For a library to be tested by this tester, the library must provide a CLI interface that conforms to the contract enforced by the tester:

```
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
2. Create a Runscope bucket. Note the Runscope bucket ID and Runscope token.
3. Add the Runscope bucket as a Segment webhook destination.

![tester](https://cldup.com/luiNQxqYu9.png)
