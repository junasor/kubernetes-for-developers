FROM mcr.microsoft.com/dotnet/sdk:8.0 AS buildstage

# Copy everything
COPY . /App
WORKDIR /App
# Restore as distinct layers

RUN dotnet restore
# Build and publish a release
RUN dotnet publish -c Release -o out

# Build runtime image
FROM mcr.microsoft.com/dotnet/runtime:8.0
WORKDIR /App
COPY --from=buildstage /App/out .
#ENTRYPOINT ["./DotNet.Docker"]
ENTRYPOINT ["dotnet", "Hello.dll"]
