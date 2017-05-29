# logrus location hook

    log := logrus.New()
    log.Hooks.Add(new(LocationHook))
    log.Info("Test location")
