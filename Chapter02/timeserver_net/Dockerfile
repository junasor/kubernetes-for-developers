FROM mcr.microsoft.com/dotnet/sdk:8.0 AS buildstage

# Copy everything
COPY . /App
WORKDIR /App
# Restore as distinct layers

RUN dotnet restore
# Build and publish a release
RUN dotnet publish -c Release -o out --no-restore

# Build runtime image
FROM mcr.microsoft.com/dotnet/aspnet:8.0
WORKDIR /App
COPY --from=buildstage /App/out .
ENTRYPOINT ["dotnet", "timeserver.dll"]
